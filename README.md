# browser-sketchbook
A frugal sketchbook for web-based Creative Coding

- dependency-free
  - browser-sketchbook follows a radical lowtech paradigm, questioning the necessity of external libraries for Creative Coding
- web-based
  - browser-sketchbook comes with a simple dedicated CLI for creating and serving sketches written in Go.

---

## Purpose

> (Creative Coding) is a process, based on exploration, iteration, reflection and discovery, where code is used as the primary medium to create a wide range of media artifacts. (Mark Mitchell, Oliver C. Bown)

This tiny framework aims to support creative programmers and programming creatives to have a fluent and frictionless workflow on writing frugal HTML, CSS and Javascript with zero dependencies. 

---

## CLI

### create.go
creates a new sketch from a specific archetype. It takes two arguments: 
1. The archetype (select one from the list)
2. the name

### build.go
scans though the sketches folder and build db.json

## Serving
- VSCode 
  - Use the popular **live-server** extension for serving the project

