package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnonymizeIDstring(t *testing.T) {
	idString := "token=eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICItOGR4MFQ2Zl9icm5XWE9TcldHRnVkdUhlYnpZQzluT0RXUUxid1RVQ0FFIn0.eyJqdGkiOiJjZjQ0YWYyNC1mZDkxLTRhYmUtYTcxOC1kMzI3ZGM2ZWViMDgiLCJleHAiOjE1NTMyMzAzMzksIm5iZiI6MCwiaWF0IjoxNTUzMTk0MzM5LCJpc3MiOiJodHRwOi8vNzAuMC4wLjgwOjgwODAvYXV0aC9yZWFsbXMvbWFzdGVyIiwic3ViIjoiMzExZjVjMmMtNzhhZS00MTNhLWE4YjktOGI3YzYyOTE3Y2Y2IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoicG9ydHdvcngiLCJhdXRoX3RpbWUiOjAsInNlc3Npb25fc3RhdGUiOiIzOTdiM2ExNi1jODAyLTQ4ZGQtODAyMi1kZjUzYWE1MTFjNjciLCJhY3IiOiIxIiwicmVzb3VyY2VfYWNjZXNzIjp7InBvcnR3b3J4Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50Iiwic3lzdGVtLmFkbWluIiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJzeXN0ZW0uYWRtaW4iLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdLCJuYW1lIjoiQmhhdmFuYSBSYWdodXBhdGhpIiwiZ3JvdXBzIjpbIioiXSwicHJlZmVycmVkX3VzZXJuYW1lIjoiYWRtaW4iLCJnaXZlbl9uYW1lIjoiQmhhdmFuYSIsImZhbWlseV9uYW1lIjoiUmFnaHVwYXRoaSIsImVtYWlsIjoiYmhhdmFuYUBnbWFpbC5jb20ifQ.hAb38WdHqw8jRENjCOLMtRQBi1q2PkOU7lJzDrqUrKWlLErcN_W-9KyAoilYZ7DEu5-SRo-YToOgBcEjJeKgevVebuP9ixUFXPOR2lymk5UBXnd5hp4G7FQtcU_eGTvfX7LtaueadK5Jd8XlBnrxi-DdTZBFGRorhn3fw6uAVGnKUEyMQ3X5Emhxsem2KnGgb-8amSnlr1wJfmoIriG7RvoODwR3y2KO1tHjo5gMmQ7XeuIZe_5IOKHjp6mU7Jqqfti2LfMNBbYL0kqBD1cctDID46Lzkt4BO78SgW1rxtuwOk9dZOU3GtuIEaOlLtm6QOV-o6Phj7FzvPvYP-uPMA,name=vContainer3"
	require.Equal(t, regxAnonString(idString, anonIDRegxes), `token=********,name=vContainer3`, "ID token anonymize failed.")

	idString = "token=lksjdkljsd:name=vContainer3"
	require.Equal(t, regxAnonString(idString, anonIDRegxes), `token=lksjdkljsd:name=vContainer3`, "ID token anonymize failed.")

	idString = "name=Cassandra-2;token=eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICItOGR4MFQ2Zl9icm5XWE9TcldHRnVkdUhlYnpZQzluT0RXUUxid1RVQ0FFIn0.eyJqdGkiOiJhMTU2MGE2Ny02NWM4LTRkODgtOGU2MS0yMzU4NjhiZDFlMmEiLCJleHAiOjE1NTQyODczMzcsIm5iZiI6MCwiaWF0IjoxNTU0MTE0NTM3LCJpc3MiOiJodHRwOi8vNzAuMC4wLjgwOjgwODAvYXV0aC9yZWFsbXMvbWFzdGVyIiwic3ViIjoiMzExZjVjMmMtNzhhZS00MTNhLWE4YjktOGI3YzYyOTE3Y2Y2IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoicG9ydHdvcngiLCJhdXRoX3RpbWUiOjAsInNlc3Npb25fc3RhdGUiOiI0MDBjNjdlMC03YjYyLTQ2OTMtODJjMS01YzU4OWE2NGIzNjAiLCJhY3IiOiIxIiwicmVzb3VyY2VfYWNjZXNzIjp7InBvcnR3b3J4Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50Iiwic3lzdGVtLmFkbWluIiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJzeXN0ZW0uYWRtaW4iLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdLCJuYW1lIjoiQmhhdmFuYSBSYWdodXBhdGhpIiwiZ3JvdXBzIjpbIioiXSwicHJlZmVycmVkX3VzZXJuYW1lIjoiYWRtaW4iLCJnaXZlbl9uYW1lIjoiQmhhdmFuYSIsImZhbWlseV9uYW1lIjoiUmFnaHVwYXRoaSIsImVtYWlsIjoiYmhhdmFuYUBnbWFpbC5jb20ifQ.JndurDBATlQ_MUZ3K-JJoQEAol1wmOqjf8Oqy-ynpxr5Q3A9uf-405j0Rl6zEx0NC45YYGUF1zgAy0Ncaa-IopI5y0ueV9G9OKbJfebC-sdJn17VN864KVMYdPwb-5WPU48oRt8O6Er5iku5o9MYCls1As64-tIMY7isACiDO_-bA74gMQ3b7vkWj0ChvR67t4WKG8dsm2GF2RyC7xEKj__0UqsORHAm1ZjWaZV7PStp_nP6i2JYnBcyN0F1iBq8nWl2WHZaNH1s6O9URD6Fvpj2kDhbkBmeSR6AvKqMtoeiNbDQASHW6ECk9mWGloP4ihnqOSjryXZRFTdNVyY3OQ"
	require.Equal(t, regxAnonString(idString, anonIDRegxes), `name=Cassandra-2;token=********`, "ID token anonymize failed.")

	idString = "name=Test;token=eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICItOGR4MFQ2Zl9icm5XWE9TcldHRnVkdUhlYnpZQzluT0RXUUxid1RVQ0FFIn0.eyJqdGkiOiJhMTU2MGE2Ny02NWM4LTRkODgtOGU2MS0yMzU4NjhiZDFlMmEiLCJleHAiOjE1NTQyODczMzcsIm5iZiI6MCwiaWF0IjoxNTU0MTE0NTM3LCJpc3MiOiJodHRwOi8vNzAuMC4wLjgwOjgwODAvYXV0aC9yZWFsbXMvbWFzdGVyIiwic3ViIjoiMzExZjVjMmMtNzhhZS00MTNhLWE4YjktOGI3YzYyOTE3Y2Y2IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoicG9ydHdvcngiLCJhdXRoX3RpbWUiOjAsInNlc3Npb25fc3RhdGUiOiI0MDBjNjdlMC03YjYyLTQ2OTMtODJjMS01YzU4OWE2NGIzNjAiLCJhY3IiOiIxIiwicmVzb3VyY2VfYWNjZXNzIjp7InBvcnR3b3J4Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50Iiwic3lzdGVtLmFkbWluIiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJzeXN0ZW0uYWRtaW4iLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdLCJuYW1lIjoiQmhhdmFuYSBSYWdodXBhdGhpIiwiZ3JvdXBzIjpbIioiXSwicHJlZmVycmVkX3VzZXJuYW1lIjoiYWRtaW4iLCJnaXZlbl9uYW1lIjoiQmhhdmFuYSIsImZhbWlseV9uYW1lIjoiUmFnaHVwYXRoaSIsImVtYWlsIjoiYmhhdmFuYUBnbWFpbC5jb20ifQ.JndurDBATlQ_MUZ3K-JJoQEAol1wmOqjf8Oqy-ynpxr5Q3A9uf-405j0Rl6zEx0NC45YYGUF1zgAy0Ncaa-IopI5y0ueV9G9OKbJfebC-sdJn17VN864KVMYdPwb-5WPU48oRt8O6Er5iku5o9MYCls1As64-tIMY7isACiDO_-bA74gMQ3b7vkWj0ChvR67t4WKG8dsm2GF2RyC7xEKj__0UqsORHAm1ZjWaZV7PStp_nP6i2JYnBcyN0F1iBq8nWl2WHZaNH1s6O9URD6Fvpj2kDhbkBmeSR6AvKqMtoeiNbDQASHW6ECk9mWGloP4ihnqOSjryXZRFTdNVyY3OQ,extra=tst"
	require.Equal(t, regxAnonString(idString, anonIDRegxes), `name=Test;token=********,extra=tst`, "ID token anonymize failed.")
}