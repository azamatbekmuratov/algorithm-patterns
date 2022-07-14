import java.util.ArrayList;
import java.util.List;

class EncodeAndDecodeStrings {
    public String encode(List<String> strs) {
        StringBuilder encodedString = new StringBuilder();
        for (String str : strs) {
            encodedString.append(str.length())
                .append("#")
                .append(str);
        }
        return encodedString.toString();
    }

    public List<String> decode(String str) {
        // Input: "4#love5#coding"
        List<String> list = new ArrayList<>();
        int i = 0;
        while (i < str.length()) {
            int j = i;
            while (str.charAt(j) != '#') {
                j++; // j=1
            }
            int length = Integer.valueOf(str.substring(i, j)); // i=0; j=1; length = 4;
            i = j + 1 + length; // i= 1 + 1 + 4 = 6
            list.add(str.substring(j + 1, i));
        }
        return list;
    }

    public static void main(String[] args) {

    }
}
