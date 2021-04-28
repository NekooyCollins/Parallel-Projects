FLAGS=-O3 -g -lpthread

RM=rm -f

EXEC=mandelbrot

all: $(EXEC)

$(EXEC):
	$(CXX) $(FLAGS) $(EXEC).cpp -c -o $(EXEC).o
	$(CXX) $(FLAGS) $(EXEC).o -o $(EXEC)

clean:
	$(RM) $(EXEC).o $(EXEC)
