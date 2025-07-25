// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureDevOps.Inputs
{

    public sealed class GetUsersFeaturesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Number of workers to process user data concurrently.
        /// 
        /// &gt; **Note** Setting `concurrent_workers` to a value greater than 1 can greatly decrease the time it takes to read the data source.
        /// </summary>
        [Input("concurrentWorkers")]
        public int? ConcurrentWorkers { get; set; }

        public GetUsersFeaturesArgs()
        {
        }
        public static new GetUsersFeaturesArgs Empty => new GetUsersFeaturesArgs();
    }
}
