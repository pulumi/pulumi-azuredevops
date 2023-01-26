import pulumi
import pulumi_random as random
import pulumi_azuredevops as ado

project_name = random.RandomId("demo-project-name", byte_length=8, prefix='demo-project-')

project = ado.Project("demo-project", name=project_name.hex)

pulumi.export("project_name", project.name)
