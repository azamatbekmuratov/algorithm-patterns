import java.util.Collections;
import java.util.HashMap;


class LongestRepeatingCharacterReplacement {
    public static int characterReplacement(String s, int k) {
        HashMap<Character, Integer> count = new HashMap<Character, Integer>();
        int res = 0;
        int l = 0;
        
        // right pointer goes till the end
        for (int r = 0; r < s.length(); r++) {
            // put character in count map
            if (count.containsKey(s.charAt(r))) {
                count.put(s.charAt(r), (1 + count.get(s.charAt(r))));
            }
            else {
                count.put(s.charAt(r), 1);
            }
            

            while(((r - l + 1) - Collections.max(count.values())) > k) {
                // decrease the character counter if we exceed k input
                count.put(s.charAt(l), (count.get(s.charAt(l)) - 1));
                l += 1;
            }
            res = Math.max(res, r - l + 1);

        }

        return res;
    }

    public static void main(String[] args) {
        int longestRepeatingSubstr = LongestRepeatingCharacterReplacement.characterReplacement("abcabcbb", 2);
        System.out.println(longestRepeatingSubstr);

    }
}