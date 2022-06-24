# About

This is a simple application in which you can set the address for the redirect through a web interface.


Then you just need to specify the address in the Android TV player (e.g. MXPlayer) ``/api/redirect`` and the service will redirect this player to the stream link.


You can use it to quickly share a long url for Android TV.

## ENV

You can also use the STREAM_URL environment variable for the initial startup. If there is no value in this environment variable, the service defaults to redirect to google.com.