public class Main {
    public static int searchInsert(int[] nums, int target) {
        int low = 0;

        int high = nums.length - 1;

        // Binary search
        while (low <= high) {
            int mid = (low + high) / 2;

            if (nums[mid] == target) {
                return mid;
            } else if (nums[mid] < target) {
                low = mid + 1;
            } else {
                high = mid - 1;
            }
        }

        return low;
    }

    public static void main(String[] args) {
        int[] numbers = {1, 3, 5, 6};

        int target = 2;

        int index = searchInsert(numbers, target);

        System.out.println("Target sonning indeksi yoki qo'shiladigan joyi: " + index);
    }
}
