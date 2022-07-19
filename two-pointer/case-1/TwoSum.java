class TwoSum {
    public int[] twoSum(int[] numbers, int target) {
        int l = 0;
        int r = numbers.length - 1;
        
        while(l < r) {
            if(numbers[l] + numbers[r] == target) break;

            if(numbers[l] + numbers[r] < target) {
                l++;
                continue;
            }
            r--;
        }

        return new int[]{l+1, r+1};
    }
    
    public static void main(String[] args) {

    }
}
