#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Output fragment color
out vec4 finalColor;

// Custom var

#define MAX_LIGHTS      3
#define MAX_SPHERES     5
#define PI              3.1415926535897932384626433832795
#define TAU             2 * PI
#define DEG_TO_RAD      TAU / 360
#define MAX_INF         1000000
#define MAX_BOUNCES     6

#define L_AMBIENT       0
#define L_POINT         1
#define L_DIRECTIONAL   2

struct Light{
    int     type;
    float   intensity;
    vec3    position;
    vec3    direction;
};

struct Sphere{
    int     specular;
    float   radius;
    float   reflective;
    float   opacity;
    float   refractionIndex;
    vec3    center;
    vec3    color;
};

struct Camera{
    mat4    rotation;
    vec3    position;
    vec3    fov;
};

struct ClosestResult {
    Sphere  closest_sphere;
    float   closest_t;
};

struct RayHitResult {
    Sphere hit;
    vec3 color;
    vec3 point;
    vec3 normal;
    vec3 objToCam;
};

struct RefResult {
    vec3 color;
    vec3 refColor;
    float r;
};

uniform vec2 res;

uniform Camera camera;
uniform Sphere spheres[MAX_SPHERES];
uniform Light lights[MAX_LIGHTS];
uniform vec3 backgroundColor;
uniform int maxBounces;

vec3 canvasToView( vec2 coord ) 
{
    return vec3(coord.xy * camera.fov.xy / res, camera.fov.z);
}

vec2 indexToCoord( vec2 indexCood )
{
    return indexCood - (res / 2);
}

vec2 intersectRaySphere( vec3 origin, vec3 direction, Sphere sphere )
{
    float radius = sphere.radius;
    vec3 edge = origin - sphere.center;

    float a = dot(direction, direction);
    float b = 2*dot(edge, direction);
    float c = dot(edge, edge) - radius * radius;

    float discriminant = b*b - 4*a*c;

    if (discriminant < 0) 
    {
        return vec2(MAX_INF, MAX_INF);
    }

    float t1 = (-b + sqrt(discriminant)) / (2*a);
    float t2 = (-b - sqrt(discriminant)) / (2*a);

    return vec2(t1,t2);
}

ClosestResult closesIntersection( vec3 origin, vec3 direction, float t_min, float t_max )
{
    float closest_t = MAX_INF;
    Sphere closest_sphere = Sphere(0, 0.0, 0.0, 0.0, 0.0, vec3(0.0), vec3(0.0));

    for(int i = 0; i < MAX_SPHERES; i++)
    {
        vec2 t = intersectRaySphere(origin, direction, spheres[i]);
        if ( t.x < closest_t && t_min < t.x && t.x < t_max ) {
            closest_t = t.x;
            closest_sphere = spheres[i];
        }
        if ( t.y < closest_t && t_min < t.y && t.y < t_max ) {
            closest_t = t.y;
            closest_sphere = spheres[i];
        }
    }
    return ClosestResult(closest_sphere, closest_t);
}

vec3 reflectRay( vec3 ray, vec3 normal )
{
    float r_dot_n = dot(ray, normal);
    return 2*normal*r_dot_n - ray;
}

float computeLight( vec3 point, vec3 normal, vec3 objToCam, int specular  )
{
    float intensity = 0;

    for (int i = 0; i < MAX_LIGHTS; i++) {
        if (lights[i].type == L_AMBIENT) {
			intensity += lights[i].intensity;
            continue;
		} 

        vec3 L = vec3(0.0);
        if (lights[i].type == L_POINT ){
            L = lights[i].position - point;
        } else {
            L = lights[i].direction;
        }

        // Shadow
        ClosestResult shadow = closesIntersection(point, L, 0.001, MAX_INF);
        if (shadow.closest_sphere.radius > 0.0 ){
            continue;
        }

        // Deffuse
        float n_dot_l = dot(normal, L);
        if (n_dot_l > 0 ){
            float length_normal = length(normal);
            float length_L = length(L);
            intensity += lights[i].intensity * n_dot_l / (length_normal * length_L);
        }

        // Specular
        if (specular > 0 ){
            vec3 reflected = reflectRay(L, normal);
            float r_dot_oc = dot(reflected, objToCam);
            if (r_dot_oc > 0.0 ){
                float length_reflected = length(reflected);
                float length_objToCam = length(objToCam);
                intensity += lights[i].intensity * pow(r_dot_oc/(length_reflected*length_objToCam), float(specular));
            }
        }		
    }

    if (intensity > 1.0) {
		intensity = 1.0;
	}

    return intensity;
}

// Find angle between two vectors
// angle = acos(angle) = (u dot v) / (length u * length v)
float rayAngleFromNormal( vec3 ray, vec3 normal)
{
    float r_dot_n = dot(ray, normal);
    float lenRay = length(ray);
    float lenNormal = length(normal);

    float angleRay = acos(r_dot_n/(lenRay*lenNormal));
    return angleRay;
}

// Formula for rotation around a arbtrary orthognal vector(both vector must be normalized)
// u = ortho vector = cross product of ray and normal
// x = vector to rotate = ray
// angle = radian
// newx = u * (u dot x) + cos(angle) * (u cross x) cross u + sin(angle)*(u cross x)
vec3 refractRay( vec3 ray, vec3 normal, float angleRay, float refractionIndex )
{
    float angleIndex = asin(sin(angleRay)/refractionIndex);
    vec3 crossRayNormal = normalize(cross(normal, ray));
    vec3 crossRayCross = cross(crossRayNormal, ray);

    vec3 c1 = crossRayNormal * dot(crossRayNormal, ray);
    vec3 c2 = cross(crossRayCross * cos(angleIndex), crossRayNormal);
    vec3 c3 = crossRayCross * sin(angleIndex);

    return normalize(c1+c2+c3);
}

RayHitResult traceRay( vec3 origin, vec3 ray, float t_min, float t_max ) 
{
    ClosestResult result = closesIntersection(origin, ray, t_min, t_max);

    if (result.closest_sphere.radius == 0) {
        return RayHitResult(
            result.closest_sphere,
            backgroundColor, 
            vec3(0.0),
            vec3(0.0),
            vec3(0.0)
        );
    }

    vec3 point = origin + result.closest_t*ray;
    vec3 normal = normalize(point - result.closest_sphere.center);
    vec3 objToCam = ray * -1;

    float lightIntensity = computeLight(point, normal, objToCam, result.closest_sphere.specular);
    vec3 localColor = result.closest_sphere.color * lightIntensity;

    return RayHitResult(
        result.closest_sphere,
        localColor,
        point,
        normal,
        objToCam
    );
}

vec3 calculateReflection( RayHitResult ray, float t_min, float t_max )
{
    vec3 colorAcc = vec3(0.0);
    RefResult ref[MAX_BOUNCES];
    int countRays = 0; // Since array is initialized early, don't overshoot on less bounces
    for ( int i = 0; i < maxBounces; i++ ){
        float r = ray.hit.reflective;
        vec3 reflected = reflectRay(ray.objToCam, ray.normal);
        
        RayHitResult rayFlected = traceRay(
            ray.point, reflected, 0.001, t_max);

        ref[i] = RefResult(
            ray.color,
            rayFlected.color,
            r
        );
        
        countRays++;

        if ( rayFlected.hit.radius <= 0.0 || 
            rayFlected.hit.reflective <= 0.0 ) {
            break;
        } 
        ray = rayFlected;
    }

    for (int i = countRays-1; i >= 0; i--) {
        if (countRays-1 == i) {
            colorAcc = ref[i].color * (1.0-ref[i].r) + ref[i].refColor * ref[i].r;
            continue;
        }
        colorAcc = ref[i].color * (1.0-ref[i].r) + colorAcc * ref[i].r;
    }
    return colorAcc;
}

vec3 calculateRefraction( RayHitResult hit, vec3 ray, float t_min, float t_max )
{
    vec3 colorAcc = vec3(0.0);
    RefResult ref[MAX_BOUNCES];
    int countRays = 0; // Since array is initialized early, don't overshoot on less bounces
    for ( int i = 0; i < maxBounces; i++ ){
        float o = hit.hit.opacity;
        float angleRay = rayAngleFromNormal(ray, hit.normal);
		vec3 refracted = refractRay(ray, hit.normal, angleRay, hit.hit.refractionIndex);
        
        RayHitResult rayFracted = traceRay(
            hit.point, refracted, 0.001, t_max);

        ref[i] = RefResult(
            hit.color,
            rayFracted.color,
            o
        );
        
        countRays++;

        if ( rayFracted.hit.radius <= 0.0 || 
            rayFracted.hit.reflective <= 0.0 ) {
            break;
        } 
        hit = rayFracted;
    }

    for (int i = countRays-1; i >= 0; i--) {
        if (countRays-1 == i) {
            colorAcc = ref[i].color * (1.0-ref[i].r) + ref[i].refColor * ref[i].r;
            continue;
        }
        colorAcc = ref[i].color * (1.0-ref[i].r) + colorAcc * ref[i].r;
    }
    return colorAcc;
}

void main()
{
    vec2 index = indexToCoord(gl_FragCoord.xy);
    vec3 direction = canvasToView(index);
    direction = (camera.rotation * vec4(direction, 0.0)).xyz;
    
    float t_min = 1;
    float t_max = MAX_INF;

    vec3 colorAcc = vec3(0.0);
    RayHitResult ray = traceRay(camera.position, direction, t_min, t_max);

    if ( ray.hit.opacity > 0.0 ) {
        ray.color = calculateRefraction(ray, direction, t_min, t_max);
    }

    if ( ray.hit.reflective > 0.0 ) {
        ray.color = calculateReflection(ray, t_min, t_max);
    }

    colorAcc = ray.color;
    finalColor = vec4(colorAcc, 1.0);
}
