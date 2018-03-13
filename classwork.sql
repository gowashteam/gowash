INSERT INTO ap.invoices
VALUES (DEFAULT, 32,'AX-014-027',DATE'8/1/2014','$434.54','&0.00','&0.00',2,DATE'8/31/2014'NULL)

UPDATE ap.invoices
SET credit_total = invoice_total * 0.1, payment_total = invoice_total - credit_total
WHERE invoice_id = 115;

DELETE FROM ap.invoices
WHERE invoice_id = 115;

SELECT hotelName FROM ap.Hotel

SELECT guestName, guestAddress FROM ap.Guest
WHERE guestAddress LIKE %London% --
ORDER BY guestName ASC;

SELECT * FROM ap.Room
WHERE type IN(double,family) AND price < 100
ORDER BY price ASC;

SELECT AVG(price) FROM ap.Room;

SELECT hotelNo, COUNT(roomNo) AS count FROM ap.Room
GROUP BY hotelNo
