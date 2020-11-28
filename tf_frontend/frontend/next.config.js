const APIURL_ENV = process.env.APIURL_ENV;
const COOKIE_DOMAIN = process.env.COOKIE_DOMAIN;

module.exports = {
  env: {
    apiUri: APIURL_ENV,
    cookieDomain: COOKIE_DOMAIN,
  },
};
