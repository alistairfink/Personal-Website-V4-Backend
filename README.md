# Personal Website Backend - V4

This is a remake of the node back end located here: https://github.com/alistairfink/Personal-Website-V3

Looking back at my previous node back end there were several major issues I was not happy about. These involved the extremely loose use of REST guidelines, the poor organization of the project, the overall use of a weakly typed language, and the performance of the system.

This lead me to a complete rewrite of the the node backend using GO which resulted in a much cleaner and maintainable architecture, faster performance, strong typing, and close following of REST and HTTP guidelines. 

NOTE: This back end is not being run right now. As of Feb 21, 2019 I am still running the Node back end in production. Why? This is because I plan on migrating all my services and database to a docker based system. This is for better organization devops wise. The only problem with that is that I am running on a OpenVZ vps. This means I cannot use docker until I am able to source a reliable KVM vps. 

