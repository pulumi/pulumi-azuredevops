import * as ado from "@pulumi/azuredevops";
import * as random from "@pulumi/random";

const randomProjectName = new random.RandomPet("demo-project-name");

const project = new ado.Project("demo-project", {
  name: randomProjectName.id,
});

export const myProjectName = project.name;
