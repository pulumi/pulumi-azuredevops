import * as ado from "@pulumi/azuredevops";
import * as random from "@pulumi/random";

const entropy = new random.RandomId("entropy", {byteLength: 8});

const randomProjectName = new random.RandomPet("demo-project-name", {keepers: {entropy: entropy.hex}});

const project = new ado.Project("demo-project", {
  name: randomProjectName.id,
});

export const myProjectName = project.name;
