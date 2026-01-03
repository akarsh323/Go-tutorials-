#!/usr/bin/env python3
"""
Improved simplification - More aggressive removal of verbose content.
"""

import os
import re
from pathlib import Path

def aggressive_simplify(file_path):
    """More aggressively simplify Go files."""
    
    with open(file_path, 'r') as f:
        content = f.read()
    
    # Step 1: Change package to intermediate
    content = re.sub(r'^package main', 'package intermediate', content, flags=re.MULTILINE)
    
    # Step 2: Remove ALL comments blocks that are >5 lines
    lines = content.split('\n')
    new_lines = []
    i = 0
    
    while i < len(lines):
        line = lines[i]
        
        # Remove large comment blocks
        if line.strip().startswith('/*'):
            # Find closing */
            comment_lines = [line]
            i += 1
            while i < len(lines):
                comment_lines.append(lines[i])
                if '*/' in lines[i]:
                    break
                i += 1
            
            # Only keep if < 4 lines (short comments)
            if len(comment_lines) < 4:
                new_lines.extend(comment_lines)
            
            i += 1
            continue
        
        # Skip lines with decorative separators
        if re.match(r'^[/=\s_-]{30,}', line):
            i += 1
            continue
        
        # Skip example function headers that are verbose
        if 'Example' in line and 'func ' not in line:
            i += 1
            continue
        
        new_lines.append(line)
        i += 1
    
    content = '\n'.join(new_lines)
    
    # Step 3: Remove KEY TAKEAWAYS and everything after main() ends
    # Find main() and keep only what's inside
    main_match = re.search(r'func main\(\) \{', content)
    if main_match:
        # Find the closing brace of main
        start = main_match.start()
        brace_count = 0
        pos = main_match.end()
        found_end = False
        
        while pos < len(content):
            if content[pos] == '{':
                brace_count += 1
            elif content[pos] == '}':
                if brace_count == 0:
                    found_end = True
                    # Keep everything up to and including this closing brace
                    content = content[:pos+1]
                    break
                brace_count -= 1
            pos += 1
    
    # Step 4: Remove multiple consecutive blank lines
    content = re.sub(r'\n{3,}', '\n\n', content)
    
    # Step 5: Clean up remaining decorative characters
    content = re.sub(r'[╔╚║╝═╞╡╤╪╟╢═]{2,}', '', content)
    
    with open(file_path, 'w') as f:
        f.write(content)
    
    return True

def main():
    """Process all detailed Go files."""
    
    base_path = Path('/Users/akarsh/GOTUT/intermediate_topics')
    detailed_files = sorted(base_path.glob('*_detailed.go'))
    
    print(f"Re-simplifying {len(detailed_files)} files with improved method...")
    
    for file_path in detailed_files:
        try:
            aggressive_simplify(str(file_path))
            print(f"✓ Improved: {file_path.name}")
        except Exception as e:
            print(f"✗ Error: {file_path.name}: {e}")
    
    print(f"\n✓ Complete!")

if __name__ == '__main__':
    main()
