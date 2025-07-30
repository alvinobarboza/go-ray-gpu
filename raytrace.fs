#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Output fragment color
out vec4 finalColor;

// Custom var

uniform vec2 res;

void main()
{
    finalColor = vec4(
        (gl_FragCoord.x*gl_FragCoord.y)/(res.x*res.y), 
        gl_FragCoord.y/res.y, 
        gl_FragCoord.x/res.x, 
        1.0
    );
}
