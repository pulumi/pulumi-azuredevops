// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class CheckRequiredTemplateRequiredTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the repository storing the template.
        /// </summary>
        [Input("repositoryName", required: true)]
        public Input<string> RepositoryName { get; set; } = null!;

        /// <summary>
        /// The branch in which the template will be referenced.
        /// </summary>
        [Input("repositoryRef", required: true)]
        public Input<string> RepositoryRef { get; set; } = null!;

        /// <summary>
        /// The type of the repository storing the template. Possible values are: `azuregit`, `github`, `githubenterprise`, `bitbucket`. Defaults to `azuregit`.
        /// </summary>
        [Input("repositoryType")]
        public Input<string>? RepositoryType { get; set; }

        /// <summary>
        /// The path to the template yaml.
        /// </summary>
        [Input("templatePath", required: true)]
        public Input<string> TemplatePath { get; set; } = null!;

        public CheckRequiredTemplateRequiredTemplateArgs()
        {
        }
        public static new CheckRequiredTemplateRequiredTemplateArgs Empty => new CheckRequiredTemplateRequiredTemplateArgs();
    }
}
