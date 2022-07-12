

class BestTimeToBuyAndSell {
    public static int maxProfit(int[] prices) {
        int maxP = 0;
        int left = 0; // left=buy
        int right = 1; // right=sell

        while(right < prices.length){
            //profitable?
            if (prices[left] < prices[right]) {
                int profit = prices[right] - prices[left];
                maxP = Math.max(maxP, profit);
            }
            else {
                left = right;
            }
            right += 1;
        }


        return maxP;
    }
    

    public static void main(String[] args) {
        int result = BestTimeToBuyAndSell.maxProfit(new int[] { 7,1,5,3,6,4 });
        System.out.println("Max profit is: " + result);
    }
}
