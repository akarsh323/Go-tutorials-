#!/usr/bin/env python3
"""
Simplify Go teaching files to match GoBootcamp pragmatic style.
Removes verbose headers and restructures for clarity.
"""

import os
import re
from pathlib import Path

def simplify_go_file(file_path):
    """Simplify a Go file by removing verbose headers and restructuring."""
    
    with open(file_path, 'r') as f:
        content = f.read()
    
    original_content = content
    
    # Step 1: Change package declaration to intermediate
    content = re.sub(r'^package main', 'package intermediate', content, flags=re.MULTILINE)
    
    # Step 2: Remove verbose multi-line comment blocks (keep simple ones)
    # Match /* ... */ blocks with many lines
    content = re.sub(
        r'/\*\n[^*]*DEEP DIVE[^*]*\*/',
        '',
        content,
        flags=re.DOTALL
    )
    
    # Remove elaborate section headers
    content = re.sub(
        r'// ={50,}.*?={50,}\n// [A-Z].*?\n// ={50,}.*?={50,}\n',
        '',
        content,
        flags=re.DOTALL | re.MULTILINE
    )
    
    # Step 3: Simplify example function names (remove "Example", "closure", etc suffixes)
    content = re.sub(r'func (closure|recursion|pointer|function)Example\d+\(\)', 
                     lambda m: 'func example()' if 'Example' not in m.group(0) else m.group(0),
                     content)
    
    # Step 4: Remove overly verbose fmt.Println headers for examples
    content = re.sub(
        r'fmt\.Println\("\\n=== Example \d+: [^"]+==="\)',
        '',
        content
    )
    
    # Step 5: Remove KEY TAKEAWAYS sections and trailing verbose content
    lines = content.split('\n')
    filtered_lines = []
    skip_section = False
    
    for i, line in enumerate(lines):
        if 'KEY TAKEAWAYS' in line or 'COMPREHENSIVE GUIDE' in line:
            skip_section = True
        elif skip_section and line.strip().startswith('func '):
            skip_section = False
        
        if not skip_section:
            filtered_lines.append(line)
    
    content = '\n'.join(filtered_lines)
    
    # Step 6: Clean up multiple blank lines
    content = re.sub(r'\n{3,}', '\n\n', content)
    
    # Step 7: Add simple header comment if missing
    if not content.startswith('package'):
        content = f'package intermediate\n\n{content}'
    
    # Step 8: Ensure main() function exists
    if 'func main()' not in content:
        # Try to find first function and add a main if needed
        # For now, just keep as is for files meant to be libraries
        pass
    
    # Only write if significantly changed
    if content != original_content:
        with open(file_path, 'w') as f:
            f.write(content)
        return True
    return False

def main():
    """Process all detailed Go files in intermediate_topics directory."""
    
    base_path = Path('/Users/akarsh/GOTUT/intermediate_topics')
    detailed_files = sorted(base_path.glob('*_detailed.go'))
    
    print(f"Found {len(detailed_files)} detailed files to process")
    
    updated = 0
    for file_path in detailed_files:
        try:
            if simplify_go_file(str(file_path)):
                updated += 1
                print(f"✓ Updated: {file_path.name}")
            else:
                print(f"- Skipped: {file_path.name} (minimal changes needed)")
        except Exception as e:
            print(f"✗ Error processing {file_path.name}: {e}")
    
    print(f"\n✓ Successfully updated {updated}/{len(detailed_files)} files")

if __name__ == '__main__':
    main()
