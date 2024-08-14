# Mainframe
This is an early WIP project for a controlpanel to have a central location for administration of both Windows and Linux services.
Some examples are:

1. PowerDNS
2. Exchange On-Prem
3. Apache2/Nginx

The idea is to have the central webpage connect to client applications that run a webserver and expose a REST API
This is not necessary for PowerDNS since it already comes with an API but Exchange and both Apache2 and Nginx does not
have a builtin API that can be used to create, update, delete, etc.
