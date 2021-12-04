# Find all the part1 and part2 cpp files
OBJ = $(shell find . -name "*part*.go" | sed 's/.go//g')

# Compile part1 or 2 for a given day
% : %.go
	go build -o $@ $<

# Compile all
all: $(OBJ)

# Remove all binaries
clean:
	rm */part1 */part2
