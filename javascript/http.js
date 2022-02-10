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
      log.error('Request: ' + host + method);
      log.error('Params: ' + params);
      log.error(error);
      log.error('\n');
    } else {
      log.info('Request: ' + host + method);
      log.info('Params: ' + params);
      log.info(body);
      log.info('\n');
    }
  });
};
