import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;

class LongestConsecutiveSequence {
    public static int longestConsecutive(List<Integer> nums){
        // input nums: 100, 4, 200, 2, 1, 3
        HashSet<Integer> numSet = new HashSet<>();
        for (Integer n : nums) {
            numSet.add(n);
        }
        int longest = 0;
        System.out.println(numSet.toString());

        for (int n : nums) {
            // check if its the start of a sequence
            if(!numSet.contains(n - 1)) {
                int length = 0;
            
                while (numSet.contains(n + length)){
                    length += 1;
                }
                longest = Math.max(length, longest);
            }
        }
        return longest;

    }

    public static void main(String[] args) {
        List<Integer> nums = new ArrayList<>();
        nums.add(100);
        nums.add(200);
        nums.add(4);
        nums.add(3);
        nums.add(2);
        nums.add(1);
        int result = LongestConsecutiveSequence.longestConsecutive(nums);
        System.out.println(result);
    }
}