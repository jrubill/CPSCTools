# python runExternalProg.py

from subprocess import call
# compile program (who needs make anyways?)
call(["g++", "-Wall", "exampleC.cpp", "-o", "program"])

# call program
call(["./program"])
