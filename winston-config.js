const winston = require('winston');
const MESSAGE = Symbol.for('message');

const jsonFormatter = (logEntry) => {
  const base = { '@timestamp': new Date(), 'stacktrace': logEntry['message'] };
  logEntry[MESSAGE] = JSON.stringify(base);
  return logEntry;
};

const logger = winston.createLogger();

logger.add(new winston.transports.Console({
  level: 'debug',
  format: winston.format(jsonFormatter)(),
}));

module.exports = logger;