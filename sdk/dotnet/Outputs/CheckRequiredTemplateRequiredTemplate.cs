// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Outputs
{

    [OutputType]
    public sealed class CheckRequiredTemplateRequiredTemplate
    {
        /// <summary>
        /// The name of the repository storing the template.
        /// </summary>
        public readonly string RepositoryName;
        /// <summary>
        /// The branch in which the template will be referenced.
        /// </summary>
        public readonly string RepositoryRef;
        /// <summary>
        /// The type of the repository storing the template. Valid values: `azuregit`, `github`, `bitbucket`. Defaults to `azuregit`.
        /// </summary>
        public readonly string? RepositoryType;
        /// <summary>
        /// The path to the template yaml.
        /// </summary>
        public readonly string TemplatePath;

        [OutputConstructor]
        private CheckRequiredTemplateRequiredTemplate(
            string repositoryName,

            string repositoryRef,

            string? repositoryType,

            string templatePath)
        {
            RepositoryName = repositoryName;
            RepositoryRef = repositoryRef;
            RepositoryType = repositoryType;
            TemplatePath = templatePath;
        }
    }
}
