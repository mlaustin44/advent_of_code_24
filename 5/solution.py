import itertools

def parse_input(in_file):
    rules = []
    pages = []
    lines = []
    with open(in_file, 'r') as f:
        lines = f.readlines()
    
    in_rules = True
    for line in lines:
        line.rstrip()
        if line == "\n":
            in_rules = False
        elif in_rules:
            rules.append(list(map(int, line.split("|"))))
        else:
            pages.append(list(map(int, line.split(","))))
    
    return rules, pages

def solve_part1(rules, page_sets):
    
    correct_page_sets = []
    for page_set in page_sets:
        set_correct = True
        for i, page in enumerate(page_set):
            rule_results = [1 for k in rules]
            for j, rule in enumerate(rules):
                # the page is the earlier one
                if page == rule[0]:
                    # go backwards through the pages to ensure the second page isn't in front of this page
                    for k in range(i - 1, 0, -1):
                        if rule[1] == page_set[k]:
                            rule_results[j] = 0
                # the page is the later one
                elif page == rule[1]:
                    # go backwards through the previous pages to make sure it fits
                    for k in range(i, len(page_set)):
                        if rule[0] == page_set[k]:
                            rule_results[j] = 0
            for rule_result in rule_results:
                if rule_result == 0:
                    set_correct = False
                    break
        if set_correct:
            correct_page_sets.append(page_set)
    
    total = 0
    for cps in correct_page_sets:
        total += cps[int(len(cps) / 2)]
    
    return total

def solve_part2(rules, page_sets):
    changed_page_sets = set()
    print(f"total page sets: {len(page_sets)}")
    for m, page_set in enumerate(page_sets):
        print(m)
        print("\t" + page_set)
        i = 0
        while(True):
            all_rules_passed = True
            page = page_set[i]
            for j, rule in enumerate(rules):
                # the page is the earlier one
                if page == rule[0]:
                    # go backwards through the pages to ensure the second page isn't in front of this page
                    for k in range(i - 1, 0, -1):
                        if rule[1] == page_set[k]:
                            all_rules_passed = False
                            changed_page_sets.add(m)
                            print(f"\t\tSwapping {page_set[i]}, {page_set[k]}")
                            page_set[i], page_set[k] = page_set[k], page_set[i]
                            i = 0
                # the page is the later one
                elif page == rule[1]:
                    # go backwards through the previous pages to make sure it fits
                    for k in range(i, len(page_set)):
                        if rule[0] == page_set[k]:
                            all_rules_passed = False
                            changed_page_sets.add(m)
                            print(f"\t\tSwapping {page_set[i]}, {page_set[k]}")
                            page_set[i], page_set[k] = page_set[k], page_set[i]
                            i = 0
            
            if (i == len(page_set) - 1) and (not all_rules_passed):
                print("Rechecking")
                i = 0
            elif(i == len(page_set) - 1) and (all_rules_passed):
                break
            else:
                i += 1

    total = 0
    print(changed_page_sets)
    for changed_set_idx in changed_page_sets:
        changed_set = page_sets[changed_set_idx]
        total += changed_set[int(len(changed_set) / 2)]
    return total

rules, pages = parse_input('input')
part1_result = solve_part1(rules, pages)
print(f"Part 1 result: {part1_result}")
part2_result = solve_part2(rules, pages)
print(f"Part 2 result: {part2_result}")