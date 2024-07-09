NAME			:=	bsq

SRCS			:=	vendor/piscine/atoi.go \
					vendor/piscine/error_handling.go \
					vendor/piscine/find_maxsquare.go \
					vendor/piscine/mapgen.go \
					vendor/piscine/open_read.go \
					vendor/piscine/print.go \
					vendor/piscine/split.go \
					main.go

all:	$(NAME)

$(NAME): $(SRCS)
		go build -o $(NAME)

clean:
		rm -f $(NAME)

.PHONY: all clean
