import pulumi
import pulumi_random as random
import pulumi_azuredevops as ado

project_name = random.RandomPet("demo-project-name")

project = ado.Project("demo-project",
                           name=project_name)

pulumi.export("project_name", project.name)


