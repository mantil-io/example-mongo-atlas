## About

This example shows how to use AWS Lambda deployed through Mantil with Mongo Atlas.
## Prerequisites

This example is created with Mantil. To download [Mantil CLI](https://docs.mantil.com/cli_install) on Mac or Linux use Homebrew 
```
brew tap mantil-io/mantil
brew install mantil
```
or check [direct download links](https://docs.mantil.com/cli_install#direct-download-linux-windows-and-mac).

To deploy this application you will need:
- [AWS account](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/)
- [Mongo Atlas Account](https://www.mongodb.com/cloud/atlas/register)

## Installation

To locally create a new project from this example run:
```
mantil new app --from mongo-atlas
cd app
```

## Configuration 

Before deploying your application you will need to create Mongo Atlas cluster. Detailed instructions on how to create a cluster can be found in [documentation](https://docs.atlas.mongodb.com/tutorial/create-new-cluster/).

Once your cluster is created you need to add connection URI to `config/environment.yml` as env variable for your function. Instructions on how to get this URI can be found [here](https://docs.atlas.mongodb.com/connect-to-cluster/).

```
project:
  env:
    CONNECTION_URI: # connection uri to mongo atlas
```

## Deploying an application

Note: If this is the first time you are using Mantil you will need to install Mantil Node on your AWS account. For detailed instructions please follow the [setup guide](https://docs.mantil.com/aws_detailed_setup/aws_credentials)

```
mantil aws install
```

After configuring the environment variable you can proceed with application deployment.

```
mantil deploy
```

This command will create a new stage for your project with default name `development` and deploy it to your node.

Now you can output the stage endpoint with `mantil env -u`. The API endpoint for your function will have the name of that function in the path, in our case that is `$(mantil env -u)/db`.

## Querying database

After deploying application you can start working with the mongo database.

What follows are invoke examples for creating, fetching and deleting items.

```
mantil invoke db/create -d '{"id":1,"name":"item"}'

mantil invoke db/get -d 1

mantil invoke db/delete -d 1
```

## Cleanup

To remove the created stage with all resources from your AWS account destroy it with

```
mantil stage destroy development
```

## Final thoughts

With this template you learned how to use AWS Lambda with Mongo Atlas. Check out our [documentation](https://docs.mantil.com/examples) to find more interesting templates. 

If you have any questions or comments on this concrete template or would just like to share your view on Mantil contact us at [support@mantil.com](mailto:support@mantil.com) or create an issue.
