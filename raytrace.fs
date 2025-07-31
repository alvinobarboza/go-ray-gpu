#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Output fragment color
out vec4 finalColor;

// Custom var

struct Light{
    int     type;
    float   intensity;
    vec3    position;
    vec3    direction;
};

struct Sphere{
    vec3    center;
    float   radius;
    vec3    color;
    int     specular;
    float   reflective;
	float   opacity;
	float   refractionIndex;
};

struct Camera{
    vec3    rotation;
    vec3    position;
    vec3    fov;
};

#define MAX_LIGHTS 3
#define MAX_SPHERES 5

uniform vec2 res;

uniform Camera camera;
uniform Sphere[MAX_SPHERES] spheres;
uniform Light[MAX_LIGHTS] lights;

vec3 canvasToView( vec2 coord ) 
{
    return vec3(coord.xy * camera.fov.xy / res, camera.fov.z);
}

vec2 indexToCoord( vec2 indexCood )
{
    return indexCood - (res / 2);
}

void main()
{
    vec2 index = indexToCoord(gl_FragCoord.xy);
    vec3 direction = canvasToView(index);
    finalColor = vec4(
        abs(camera.position) / 10,
        1.0
    );
}
