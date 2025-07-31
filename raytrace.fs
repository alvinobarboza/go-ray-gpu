#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Output fragment color
out vec4 finalColor;

// Custom var

#define MAX_LIGHTS 3
#define MAX_SPHERES 5
#define PI 3.1415926535897932384626433832795
#define TAU 2 * PI
#define DEG_TO_RAD TAU / 360
#define MAX_INF 1000000

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

uniform vec2 res;

uniform Camera camera;
uniform Sphere spheres[MAX_SPHERES];
uniform Light lights[MAX_LIGHTS];
uniform vec3 backgroundColor;

vec3 canvasToView( vec2 coord ) 
{
    return vec3(coord.xy * camera.fov.xy / res, camera.fov.z);
}

vec2 indexToCoord( vec2 indexCood )
{
    return indexCood - (res / 2);
}

vec2 intersectRaySphere(vec3 origin, vec3 direction, Sphere sphere)
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

void main()
{
    vec2 index = indexToCoord(gl_FragCoord.xy);
    vec3 direction = canvasToView(index);
    direction = (camera.rotation * vec4(direction, 0.0)).xyz;

    float closest_t = MAX_INF;
    Sphere closest_sphere = Sphere(0, 0.0,0.0,0.0, 0.0,vec3(0.0),vec3(0.0));

    float t_min = 1;
    float t_max = MAX_INF;

    for(int i = 0; i < MAX_SPHERES; i++)
    {
        vec2 t = intersectRaySphere(camera.position, direction, spheres[i]);
        if ( t.x < closest_t && t_min < t.x && t.x < t_max ) {
            closest_t = t.x;
            closest_sphere = spheres[i];
        }
        if ( t.y < closest_t && t_min < t.y && t.y < t_max ) {
            closest_t = t.y;
            closest_sphere = spheres[i];
        }
    }
    if (closest_sphere.radius == 0) {
        finalColor = vec4(backgroundColor, 1.0);
        return;
    }
    finalColor = vec4(closest_sphere.color,1.0);
}
