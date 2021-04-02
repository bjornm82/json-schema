package internal

// Table shape
// expect_column_to_exist *
// expect_table_columns_to_match_ordered_list
// params: column_list ["string"]
// expect_table_columns_to_match_set
// params:
// - column_set ["string", "string"]
// - exact_match: "boolean"
// expect_table_row_count_to_be_between
// - min_value: integer, null
// - max_value: integer, null
// expect_table_row_count_to_equal
// - value: integer
// expect_table_row_count_to_equal_other_table
// - other_table_name: string (HOW??)

// Missing values, unique values, and types
// expect_column_values_to_be_unique
// - column: string
// - mostly: float (0-1)
// expect_column_values_to_not_be_null
// - column: string
// - mostly: float (0-1)
// expect_column_values_to_be_null
// - column: string
// - mostly: float (0-1)
// expect_column_values_to_be_of_type ??
// expect_column_values_to_be_in_type_list ??

// Sets and ranges
// expect_column_values_to_be_in_set
// expect_column_values_to_not_be_in_set
// expect_column_values_to_be_between
// - column: string
// - min_value: integer
// - max_value: integer
// expect_column_values_to_be_increasing
// expect_column_values_to_be_decreasing

// String matching
// expect_column_value_lengths_to_be_between
// expect_column_value_lengths_to_equal
// expect_column_values_to_match_regex
// expect_column_values_to_not_match_regex
// expect_column_values_to_match_regex_list
// expect_column_values_to_not_match_regex_list
// expect_column_values_to_match_like_pattern
// expect_column_values_to_not_match_like_pattern
// expect_column_values_to_match_like_pattern_list
// expect_column_values_to_not_match_like_pattern_list

// Datetime and JSON parsing
// expect_column_values_to_match_strftime_format
// expect_column_values_to_be_dateutil_parseable
// expect_column_values_to_be_json_parseable
// expect_column_values_to_match_json_schema

// Aggregate functions
// expect_column_distinct_values_to_be_in_set
// expect_column_distinct_values_to_contain_set
// expect_column_distinct_values_to_equal_set
// expect_column_mean_to_be_between
// expect_column_median_to_be_between
// expect_column_quantile_values_to_be_between
// expect_column_stdev_to_be_between
// expect_column_unique_value_count_to_be_between
// expect_column_proportion_of_unique_values_to_be_between
// expect_column_most_common_value_to_be_in_set
// expect_column_max_to_be_between
// expect_column_min_to_be_between
// expect_column_sum_to_be_between

// Multi-column
// expect_column_pair_values_A_to_be_greater_than_B
// expect_column_pair_values_to_be_equal
// expect_column_pair_values_to_be_in_set
// expect_select_column_values_to_be_unique_within_record
// expect_multicolumn_sum_to_equal
// expect_column_pair_cramers_phi_value_to_be_less_than
// expect_compound_columns_to_be_unique

// Distributional functions
// expect_column_kl_divergence_to_be_less_than
// expect_column_bootstrapped_ks_test_p_value_to_be_greater_than
// expect_column_chisquare_test_p_value_to_be_greater_than
// expect_column_parameterized_distribution_ks_test_p_value_to_be_greater_than

// FileDataAsset
// File data assets reason at the file level, and the line level (for text data).
// expect_file_line_regex_match_count_to_be_between
// expect_file_line_regex_match_count_to_equal
// expect_file_hash_to_equal
// expect_file_size_to_be_between
// expect_file_to_exist
// expect_file_to_have_valid_table_header
// expect_file_to_be_valid_json
