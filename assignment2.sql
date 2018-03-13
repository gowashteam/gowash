//Name: Okey Onyia
//Student ID: 0883981
//1
SELECT title as bookTitle FROM ap.book

//2
SELECT * FROM ap.borrower
ORDER BY borrowerName ASC

//3
SELECT title as titles
FROM ap.book WHERE year = 2005 AND author = 'J. K. Rowling'

//4
SELECT count(ISBN) FROM ap.book
where ISBN = '0-7475-8108-8'

//5
SELECT * FROM ap.bookCopy
where available = TRUE AND ISBN =
    (SELECT ISBN FROM ap.book
    WHERE title = "Harry Potter")

//6
SELECT borrowerName FROM ap.borrower
WHERE borrowerNo = (
  SELECT borrowerNo FROM ap.bookLoan
  WHERE copyNo = (
    SELECT copyNo FROM ap.bookCopy
    WHERE ISBN = (
      SELECT ISBN FROM ap.book
      WHERE title = 'Lord of the Rings'
    )
  )
)

//7
SELECT borrowerName FROM ap.borrower
WHERE borrowerNo = (
  SELECT borrowerNo FROM ap.bookLoan
  WHERE dateDue = CURDATE();
)

//8
SELECT * FROM ap.book WHERE ISBN IN (
  SELECT ISBN FROM ap.bookCopy
  WHERE copyNo IN (
      SELECT copyNo, count(*) FROM ap.bookLoan
      ORDER BY copyNo DESC
      LIMIT 10
  )
)
