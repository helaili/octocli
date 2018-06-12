../../../../bin/octocli org list -s octodemo.com -k $GHE_AUTH_TOKEN
../../../../bin/octocli org list -l helaili -s octodemo.com -k $GHE_AUTH_TOKEN
../../../../bin/octocli team list -o OctOcheese -s octodemo.com -k $GHE_AUTH_TOKEN
../../../../bin/octocli team members list -t dev-team -o OctOcheese -s octodemo.com -k $GHE_AUTH_TOKEN
# ../../../../bin/octocli team create -t x-team -o OctOcheese -s octodemo.com -d "X Men" -k $GHE_AUTH_TOKEN
../../../../bin/octocli team repositories list -t dev-team -o OctOcheese -s octodemo.com -k $GHE_AUTH_TOKEN
../../../../bin/octocli repo list -l OctoCheese -s octodemo.com -k $GHE_AUTH_TOKEN
../../../../bin/octocli repo audit -l OctoCheese -r Calculator -s octodemo.com -k $GHE_AUTH_TOKEN
../../../../bin/octocli repo audit -l OctoCheese -r Calculator -s octodemo.com -k $GHE_AUTH_TOKEN -g -m
