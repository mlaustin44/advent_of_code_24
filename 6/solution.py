import copy
from pprint import pprint

def parse_input(input_file):
    with open(input_file, 'r') as f:
        raw_input = f.readlines()

    input = [list(l.rstrip()) for l in raw_input]
    
    return input

# use rc notation to make the arrays easier
# directions with the rc deltas
dirc = {
    'up': (-1, 0),
    'down': (1, 0),
    'right': (0, 1),
    'left': (0, -1)
}
dirc_change = {
    'up': 'right',
    'right': 'down',
    'down': 'left',
    'left': 'up'
}
dir_decode = {
    '^': 'up',
    '>': 'right',
    '<': 'left',
    'v': 'down'
}

def solve_part1(input):
    grid = copy.deepcopy(input)

    current_pos = None
    current_dir = None

    # find the guard start position
    for r in range(len(grid)):
        done = False
        for c in range(len(grid[0])):
            if grid[r][c] in ['^', '<', '>', 'v']:
                current_pos = [r, c]
                current_dir = dir_decode[grid[r][c]]
                done = True
                break
        if done:
            break
    in_grid = True
    rows = len(grid)
    cols = len(grid[0])
    while in_grid:
        # mark the current position as visited
        grid[current_pos[0]][current_pos[1]] = 'X'

        # calculate the next step
        next_r = current_pos[0] + dirc[current_dir][0]
        next_c = current_pos[1] + dirc[current_dir][1]

        # check if we've left the grid
        if (next_r < 0) or (next_r == rows) or (next_c < 0) or (next_c == cols):
            in_grid = False
            break
        
        # check for hitting a wall
        if grid[next_r][next_c] == '#':
            current_dir = dirc_change[current_dir]
        else:
            # update the grid
            current_pos = [next_r, next_c]

    # count up the visited squares
    visited = 0
    for r in range(len(grid)):
        for c in range(len(grid[0])):
            if grid[r][c] == 'X':
                visited += 1
    
    return visited


def check_for_loop(input):
    grid = input

    current_pos = None
    current_dir = None

    # find the guard start position
    for r in range(len(grid)):
        done = False
        for c in range(len(grid[0])):
            if grid[r][c] in ['^', '<', '>', 'v']:
                current_pos = [r, c]
                current_dir = dir_decode[grid[r][c]]
                done = True
                break
        if done:
            break
    in_grid = True
    rows = len(grid)
    cols = len(grid[0])
    max_positions = rows * cols
    visited_pos = 0
    while in_grid:
        # mark the current position as visited
        grid[current_pos[0]][current_pos[1]] = 'X'

        # calculate the next step
        next_r = current_pos[0] + dirc[current_dir][0]
        next_c = current_pos[1] + dirc[current_dir][1]

        # check if we've left the grid
        if (next_r < 0) or (next_r == rows) or (next_c < 0) or (next_c == cols):
            in_grid = False
            break
        
        # check for hitting a wall
        if (grid[next_r][next_c] == '#') or (grid[next_r][next_c] == 'O'):
            current_dir = dirc_change[current_dir]
        else:
            # update the grid
            current_pos = [next_r, next_c]
            visited_pos += 1

        if visited_pos > max_positions:
            return True
    
    return False


def solve_part2(input):
    rows = len(input)
    cols = len(input)
    loop_options = 0

    # horrible double for loop, but it works!
    for r in range(rows):
        for c in range(cols):
            if (input[r][c] == '.'):
                grid = copy.deepcopy(input)
                grid[r][c] = 'O'
                has_loop = check_for_loop(grid)
                if has_loop:
                    loop_options += 1

    return loop_options




input = parse_input('input')
part1_soln = solve_part1(input)
part2_soln = solve_part2(input)
print(f'Part 1 solution: {part1_soln}')
print(f'Part 2 solution: {part2_soln}')