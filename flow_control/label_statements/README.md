# Labels Statements

-   Labels are used in break, continue, and goto statements.
-   It is illegal to define a label that is never used.
-   In contrast to other identifiers, labels are not block scoped and do not     
    conflict with identifiers that are not labels. They live in another space.
-   The scope of a labels is the body of the function in which it is declared
    and excludes the body of any nested function.
-   Most of the time labels are used to terminate outer enclosing loops