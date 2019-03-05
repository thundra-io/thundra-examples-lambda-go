# Thundra Lambda Go Example Project

This is a simple example to trace your lambda go functions with [Thundra](https://www.thundra.io/).

If you haven't done already start by installing **serverless** by:
```bash
npm install -g serverless
```

1 - Paste your **API Key** you get from the [Thundra web console](https://console.thundra.io/) to the `thundra_apiKey` environment variable in the `serverless.yml` file.`

2- **Build** and **deploy** the project using the `deploy.sh`
```bash
sh deploy.sh
```
4- **Invoke** `helloworld` function in AWS Lambda Console, or with the Serverless framework running the following command in the project directory:
```
sls invoke -f helloworld
```

5- **Observe** your invocation on [Thundra web console](https://console.thundra.io/)!

  Now you have successfully integrated Thundra Go agent to your Lambda function! It's time to create multiple spans using OpenTracing interface, add logs and configure more advanced tracing capabilities of Thundra Go agent. Check out [the documentation](https://docs.thundra.io/docs/go-installation-and-configuration) and [the blog post](https://medium.com/thundra/announcing-opentracing-compatibility-for-go-agent-e805f2f2428) for more detailed information.
