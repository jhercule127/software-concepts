# Code Structure

## Bad Comments

- **Try to Avoid Comments:** Comments often indicate redundant information and can clutter the codebase. Strive to write self-explanatory code that doesn't rely heavily on comments.

- **Dividers/Block Markers:** Use dividers or block markers to indicate sections such as global variables, classes, etc., instead of relying solely on comments.

- **Misleading Comments:** Avoid comments that are misleading or no longer accurate.

- **Commented-Out Code:** Remove commented-out code rather than leaving it in the codebase. It can clutter the code and make it harder to maintain. [Example](https://github.com/academind/clean-code-course-code/blob/comments-formatting-01-bad-comments/bad-comments.ts).

## Good Comments

- **Legal Information:** Include legal information such as copyrights.

- **Explanations that Can't be Explained by Good Naming:** Use comments to explain complex logic or provide context that cannot be conveyed through naming alone.

- **Warnings and Todo Notes:** Use comments to highlight warnings or indicate areas for future improvement (e.g., TODO notes).

- **Documentation Comments:** Depending on the project, include documentation comments to provide information about functions, classes, and other elements in the code. [Example](https://github.com/academind/clean-code-course-code/blob/comments-formatting-02-good-comments/02-good-comments.ts).

## Code Formatting

Code formatting improves readability and conveys meaning effectively.

### Vertical Formatting

- **Readability Like an Essay:** Ensure code is readable like an essay without too many jumps. If a file becomes too large, consider splitting it into multiple files based on different concepts.

- **Separate Different Areas of Code:** Separate different areas of code with spacing to improve readability. Related concepts should be close to each other. [Example](https://github.com/academind/clean-code-course-code/blob/comments-formatting-03-vertical-formatting/vertical-formatting.js).

- **Do Not Separate Similar Concepts:** Similar concepts should not be separated by excessive spacing.

### Horizontal Formatting

- **Readable Without Scrolling:** Lines of code should be readable without horizontal scrolling. Use indentation and break long statements into shorter ones.

- **Clear but Not Unreadable Long Names:** Use clear and descriptive names, even if they are long, to improve code readability.

- **Follow Language-Specific Conventions:** Formatting rules may differ between languages. Follow language-specific conventions and guidelines to ensure consistency and readability.

