import * as ado from "@pulumi/azuredevops";
import * as random from "@pulumi/random";

const randomProjectName = new random.RandomId("demo-project-name", {byteLength: 8, prefix: "demo-project-"});

const project = new ado.Project("demo-project", {
    name: randomProjectName.hex,
});

export const myProjectName = project.name;
