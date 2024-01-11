// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class GetUsersFeaturesInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of workers to process user data concurrently.
        /// 
        /// &gt; **Note** Setting `concurrent_workers` to a value greater than 1 can greatly decrease the time it takes to read the data source.
        /// </summary>
        [Input("concurrentWorkers")]
        public Input<int>? ConcurrentWorkers { get; set; }

        public GetUsersFeaturesInputArgs()
        {
        }
        public static new GetUsersFeaturesInputArgs Empty => new GetUsersFeaturesInputArgs();
    }
}
