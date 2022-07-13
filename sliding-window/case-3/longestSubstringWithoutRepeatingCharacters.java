import java.util.HashSet;

class LongestSubstringWithoutRepeatingCharacters {
    public static int longestSubstring(String input) {
        HashSet<Character> charSet = new HashSet<Character>();
        int l = 0;
        int result = 0;
        
        for (int r = 0; r < input.length(); r++) {
            while (charSet.contains(input.charAt(r))){
                charSet.remove(input.charAt(l));
                l += 1;
            }
            charSet.add(input.charAt(r));
            result = Math.max(result, r-l+1);
        }
        return result;
    }

    public static void main(String[] args){
        int longestSubstr = LongestSubstringWithoutRepeatingCharacters.longestSubstring("abcabcbb");
        System.out.println(longestSubstr);

    }
}
