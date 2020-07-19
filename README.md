# Poll-to-push pattern example


Basically this pattern implemented with `Go`.

![Image](https://d2908q01vomqb2.cloudfront.net/1b6453892473a467d07372d45eb05abc2031647a/2019/08/27/Arch.png)

## Goals

I'm always learning with purpose, I made this to:

- remind myself how *step functions* work
- tackle testing with new `aws-sdk-go-v2`. Overall the experience could be better.

## How to run this 

Make sure you have [wscat](https://www.npmjs.com/package/wscat) installed.

1. In the root directory perform these in your terminal:
    - `npm install`
    - `npm run deploy`

2. Copy the `wss` endpoint from the terminal

3. Copy the `POST` endpoint from the terminal.

4. Make a post request to
    ```text
    POST_ENDPOINT/being-work
    ```

    With payload
    ```json
    {"name":  "YOUR_NAME_HERE"}
    ```

5. Connect to the websocket endpoint using `wscat`

    ```text
    wscat -c wss_ENDPOINT
    ```
6. Listen to the results

    ```json
    {"action":"sendmessage","data":{"taskID": "ID_YOU_GOT_FROM_HTTP_POST"}}
    ```


