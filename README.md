<img src="https://user-images.githubusercontent.com/74418041/206620270-75cd9cf3-63c4-4b1c-8611-a4acf05ed1c0.gif" width=50%>

> A Remote Access Trojan & Post-Exploitation Framework 

# Introduction
Nemo is yet another post-exploitation framework, focused around collaboration. Collaboration among multiple operators is managed by a "swarm", which can be remotely managed from the command and control. Payloads are *not* dynamically generated; nemo is not intended to be used in production. Do keep in mind that I'm fairly new to malware dev, and this is just one of those shower thought projects that has magically came to life.

# Installation & Usage
Nemo consists of a number of components, namely:
- [**client**](./src/service/client) - A CLI shell utility to manage payloads and communicate with operators. The main interface you will be interacting with. 
- [**server**](./src/server) - The remote command and control server which will be the main access point into the payloads. The daemon usually manages the server. 
- [**daemon**](./src/service/daemon) - The primary access point for the CLI to communicate with the server; a websocket API server. 
- [**payload**](./src/service/payload) - Payload generation scripts and templates. These are the payloads that are slightly modified during generation.

# Built With
![](https://skillicons.dev/icons?i=go,python,rust)
