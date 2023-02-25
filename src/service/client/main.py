from msvcrt import getch
import time
import sys
import os

import subprocess
from shutil import copyfile
from signal import signal, SIGINT
from typing import List
from lib.prepend_shell import prepend_interface

user_history = []


payload_help_menu = """
Usage: payload [command] [arguments]

Commands:
    generate         Dynamically generates a payload 

Options:
    -O, --output     Output filename
    -o, --obfuscate  Whether to obfuscate the payload or not.
    -h, --help       Display this help menu.
            
        """

payload_gen_help_menu = """
Usage: payload generate [ARGUMENTS]

Arguments:
    name     (REQUIRED)  Readable identifier for the payload to be generated. 
    platform (REQUIRED)  Platform the payload is generated for.

Options:
    -O, --output     Output filename
    -o, --obfuscate  Whether to obfuscate the payload or not.
    -h, --help       Display this help menu.
                """

def help_handler(args: List[any] or None = None):
    commands = {
        "help": "Display this help message.",
        "swarm": "Perform swarm related operations.",
        "payload": "Perform payload generation and other payload operations.",
    }

    commands_menu = ""

    for command, description in commands.items():
        commands_menu += f"{command} - {description}\n"

    if args is None or args == []:
        print("\n", commands_menu)
    elif args is not None or args != []:
        commands_menu = f"{args[0]} - {commands[args[0]]}\n"
        print(commands_menu)


def payload_handler(args: List[any] or None = None):
    if args is None or args == [] or args[0] == "-h" or args[0] == "help":
        print(payload_help_menu)

    if args != None and len(args) != 0:
        if args[0] == "generate":
            opt = args[1]
            if opt  == (" " or ""):
                print("Please provide the required options.")
                print(payload_gen_help_menu)
            if opt == ("-h" or "--help"):
                print(payload_gen_help_menu)
            if opt.__contains__("name"):
                api_key = input("Please enter an ngrok API key: ")
                # os.system(f"export NGROK_API_KEY={api_key}")
                # print("set", api_key)
                mod_env = os.environ.copy()
                mod_env["NGROK_API_KEY"] = api_key
                
                out = subprocess.Popen("../daemon/daemon", stdout=subprocess.PIPE, env=mod_env, bufsize=1)

                for cur in iter(out.stdout.readline(), b''):
                    print(cur)
                

                signal(SIGINT, lambda _: out.send_signal(SIGINT))
                
                
    elif args == None and len(args) == 0:
        print("")


commands_registry = {"help": help_handler, "payload": payload_handler}


def get_char():
    first_char = getch()
    if first_char == "\x1b":
        return {"[A": "up", "[B": "down", "[C": "right", "[D": "left"}[
            getch() + getch()
        ]
    else:
        return first_char


def run_command(command, args):
    commands_registry[command](args)
    argsStr = ""
    for arg in args:
        argsStr += arg

    user_history.append(f"{command} {argsStr}".strip())


def main():
    while True:
        prepend_interface()
        command = input("")

        if command != ("" or None):
            try:
                argv = command.split(" ")
                command = argv[:][0]
                argv.remove(command)
                run_command(command, argv)
            except:
                if (command == "payload") and len(argv) >= 2 and not (argv[1].__contains__("name")):
                    print("Please provide the required options.")
                    print(payload_gen_help_menu)
                elif (command == "payload"):
                    if len(argv) >= 2 and (argv[1] == ("-h" or "--help" or "" or " ")):
                        print(payload_gen_help_menu)
                    else: 
                        print("", end="")
                else:
                    print("Unknown command.")
                    help_handler()


def handle_close(_sig, _frame):
    print("\nQuitting client.")
    sys.exit(0)


def handle_history():
    ordered_user_history = user_history[::-1]
    print(ordered_user_history[len(ordered_user_history) - 1], end="")


signal(SIGINT, handle_close)
main()