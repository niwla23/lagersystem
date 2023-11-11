import confuse

config = confuse.Configuration('LagersystemOperator', __name__)
config.set_file('./config.yaml')

def reload():
    config.clear()
    config.set_file('./config.yaml')