const request = require('request');
const log = require('tracer').colorConsole({
  format: '{{message}}',
});

const host = 'http://127.0.0.1:7545/';

exports.post = (method, params) => {
  const options = {
    uri: host + method,
    method: 'POST',
    body: {
      jsonrpc: '2.0',
      method: method,
      params: params,
      id: 1,
    },
    json: true,
  };

  request.post(options, (error, response, body) => {
    if (error != null) {
      log.error(error);
    }
    log.info(body);
  });
};
