# Sample Go module

This is a simple Go module, the goal of this module is not the library itself, 
but to create an automated pipeline with GH actions to **detect breaking changes**.

## Use case
We have a library that is imported by multiple projects, and we want to be aware 
of breaking changes as soon as they are introduced. (Yeah, ideally we would follow
semantic versioning and this cool stuff, but our use case is a bit different).

## How it works
After every push to the main branch, the GH action `trigger-build.yml` will 
create a "[workflow dispatch event](https://docs.github.com/en/rest/actions/workflows?apiVersion=2022-11-28#create-a-workflow-dispatch-event)",
aka, trigger a GH action on another repository. By doing this, we can run another
workflow that will upgrade the Go module and build it with the new version. If the build
fails, we can assume that the changes are breaking.
