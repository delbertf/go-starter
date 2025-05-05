## golang starter project

this is a minimal starter project with:
* golang
* tailwind + daisyui
* templ 
* data-star

### features

- Hot-Reload for go files, css and templ changes
- No nodejs, we use the tailwindcss binary instead
- SSE with data-star and templ
- file structure is domain driven / feature oriented (better for bigger apps)


### how to start

clone this repo and rename it to your project name with

```find . -type f -exec sed -i 's/go2start/myproject/g' {} +```

