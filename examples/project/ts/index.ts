import * as ado from "@pulumi/azuredevops";
import * as random from "@pulumi/random";

const randomProjectName = new random.RandomPet("demo-project-name");

const project = new ado.Project("demo-project", {
    projectName: randomProjectName.id,
});

export const myProjectName = project.projectName;
