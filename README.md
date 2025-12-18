# Modular Go Server

A lightweight, modular server framework designed to separate server logic, from transport layer and application protocols.

## Goal

To create modular server, which will let user plug in different types of tranport protocols and application protocols.

## Roadmap

- Basic Architecture Setup (halfway done)
- TCP Transport Layer (halway done)
- UDP Transport Layer (needs to be implemented)
- Application Protocols (HTTP, etc.) (not implemented at all)

## Architecture

The project is split into `transport` (handling the bits/bytes) and `server` (handling the logic), allowing for easy extensibility. In the future am also going to add application protocols layer too.
