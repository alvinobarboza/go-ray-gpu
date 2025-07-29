#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Output fragment color
out vec4 finalColor;

// Custom var

uniform vec2 resolution;

void main()
{
    finalColor = vec4(gl_FragCoord.x/resolution.x,gl_FragCoord.x/resolution.x,gl_FragCoord.x/resolution.x,1.0);
}
