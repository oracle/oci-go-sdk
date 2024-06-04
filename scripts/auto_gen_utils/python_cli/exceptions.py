class BotException(Exception):
    """Base exception for self-service jira bit"""


class CliException(BotException):
    def __init__(self, message, line):
        super().__init__()
        self.message = message
        self.line = line
