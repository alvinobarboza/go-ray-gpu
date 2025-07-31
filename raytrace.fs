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
uniform Light[MAX_LIGHTS] ligths;

void main()
{
    finalColor = vec4(
        abs(camera.position) / 10,
        1.0
    );
}
