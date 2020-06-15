module.exports = {
  client: {
    service: {
      name: 'chat-it',
      url: 'http://localhost:8080/query',
      // optional headers
      // optional disable SSL validation check
      skipSSLValidation: true
    }
  }
};
