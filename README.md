# HeyApple

**HeyApple** is an online food & nutrient tracker, recipe manager, and shopping list generator.

There are many apps out there for tracking nutrients and just as many for managing recipes and
generating shopping lists. There are hardly any that combine these two into a full food management
solution. HeyApple fills that void.

You can track your daily nutrient intake down to individual minerals and vitamins, combine food
into recipes, store cooking instructions and prep times. Shopping lists are generated automatically
by simply selecting the days you plan to do your shopping for.

## Building & Testing

Execute the following command in the project's root directory:
```
$ make build
```
This will perform most unit tests, build the server app and the web frontend, and put everything
in the `./out` directory. All the necessary files to run and deploy the app are embedded into the
executable file, so this is the only file that needs to be distributed.
```
$ make run
```
This will run the application directly from the `./out` directory. By default, the app listens
on `http://localhost:8080` for incoming requests. Open it in a browser to start the web frontend.
```
$ make test
```
This will execute all unit tests in the project, including tests that detect race conditions.
It will take a while for all of these to complete, which is why they are not executed by default
when building the app.
