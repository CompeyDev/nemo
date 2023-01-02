import time
import sys
import os

import subprocess
from shutil import copyfile
from signal import signal, SIGINT
from lib.prependShell import prependInterface

userHistory = []


payloadHelpMenu = """
Usage: payload [command] [arguments]

Commands:
    generate         Dynamically generates a payload 

Options:
    -O, --output     Output filename
    -o, --obfuscate  Whether to obfuscate the payload or not.
    -h, --help       Display this help menu.
            
        """

payloadGenHelpMenu = """
Usage: payload generate [ARGUMENTS]

Arguments:
    name     (REQUIRED)  Readable identifier for the payload to be generated. 
    platform (REQUIRED)  Platform the payload is generated for.

Options:
    -O, --output     Output filename
    -o, --obfuscate  Whether to obfuscate the payload or not.
    -h, --help       Display this help menu.
                """

def helpHandler(args: [] or None = None):
    commands = {
        "help": "Display this help message.",
        "swarm": "Perform swarm related operations.",
        "payload": "Perform payload generation and other payload operations.",
    }

    commandsMenu = ""

    for command, description in commands.items():
        commandsMenu += f"{command} - {description}\n"

    if args is None or args == []:
        print("\n", commandsMenu)
    elif args is not None or args != []:
        commandsMenu = f"{args[0]} - {commands[args[0]]}\n"
        print(commandsMenu)


def payloadHandler(args: [] or None = None):
    if args is None or args == [] or args[0] == "-h" or args[0] == "help":
        print(payloadHelpMenu)

    if args != None and len(args) != 0:
        if args[0] == "generate":
            opt = args[1]
            if opt  == (" " or ""):
                print("Please provide the required options.")
                print(payloadGenHelpMenu)
            if opt == ("-h" or "--help"):
                print(payloadGenHelpMenu)
            if opt.__contains__("name"):
                api_key = input("Please enter an ngrok API key: ")
                os.system(f"export NGROK_API_KEY={api_key}")
                print("set", api_key)
                modEnv = os.environ.copy()
                modEnv["NGROK_API_KEY"] = api_key
                out = subprocess.Popen("../daemon/daemon", stdout=subprocess.PIPE, env=modEnv, bufsize=1)

                for cur in iter(out.stdout.readline(), b''):
                    print(cur)
                

                signal(SIGINT, lambda handler: out.send_signal(SIGINT))
                
                
    elif args == None and len(args) == 0:
        print("")


commandsRegistry = {"help": helpHandler, "payload": payloadHandler}


def getChar():
    first_char = getch()
    if first_char == "\x1b":
        return {"[A": "up", "[B": "down", "[C": "right", "[D": "left"}[
            getch() + getch()
        ]
    else:
        return first_char


def runCommand(command, args):
    commandsRegistry[command](args)
    argsStr = ""
    for arg in args:
        argsStr += arg

    userHistory.append(f"{command} {argsStr}".strip())


def main():
    while True:
        prependInterface()
        command = input("")

        if command != ("" or None):
            try:
                argv = command.split(" ")
                command = argv[:][0]
                argv.remove(command)
                runCommand(command, argv)
            except:
                if (command == "payload") and len(argv) >= 2 and not (argv[1].__contains__("name")):
                    print("Please provide the required options.")
                    print(payloadGenHelpMenu)
                elif (command == "payload"):
                    if len(argv) >= 2 and (argv[1] == ("-h" or "--help" or "" or " ")):
                        print(payloadGenHelpMenu)
                    else: 
                        print("", end="")
                else:
                    print("Unknown command.")
                    helpHandler()


def handleClose(sig, frame):
    print("\nQuitting client.")
    sys.exit(0)


def handleHistory():
    orderedUserHistory = userHistory[::-1]
    print(orderedUserHistory[len(orderedUserHistory) - 1], end="")


signal(SIGINT, handleClose)
main()