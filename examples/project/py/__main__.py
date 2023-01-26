import pulumi
import pulumi_random as random
import pulumi_azuredevops as ado

entropy = random.RandomId("entropy", byte_length=8)

project_name = random.RandomPet("demo-project-name", keepers={'entropy': entropy.hex})

project = ado.Project("demo-project",
                           name=project_name)

pulumi.export("project_name", project.name)
