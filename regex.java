// Pt. 1: exact match
//  matches("hello world", "hello") -> true
//  matches("hello world", "elo") -> false
//  matches("hello world", "elllo") -> false
//  matches("ababac", "abac") -> true
// Pt. 2: wildcard '.' - matches exactly one of any character in text
//  matches("hello world", "he.lo") -> true
//  matches("hello world", "h.o") -> false

boolean matches(String text, String query){

    if(querry.length() > text.length())
        return false;

    int count = 0;
    char[] textArray = text.toCharArray();
    char[] querryArray = query.charArray();
    boolean check = true;
    int i = 0;
    int j = 0;
    int trackTextLoop = 0;

    while(j != textArray.length()){
        if(!check){
            j = trackTextLoop++;
            trackTextLoop = j;
        }
        if((querryArray[i] == textArray[j]) || (querryArray[i] == '.')){
            //match found
            check = true;
            count++; // 3
            i++;
            j++;
            continue;

        }else{
            check = false;
            i = 0;
            count = 0;
        }


    }

    return (count == querryArray.length());
}
