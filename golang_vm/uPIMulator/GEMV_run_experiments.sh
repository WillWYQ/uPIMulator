#!/bin/bash

# Define the values to replace #
VALUES=(1 2 4 8 16 32 64 128)

# Iterate over each value
for NUM_DPUS in "${VALUES[@]}"; do
    echo "Running uPIMulator with --num_dpus_per_rank=${NUM_DPUS}"

    # Run the command with the current value
    ./build/uPIMulator \
        --verbose 2 \
        --root_dirpath /uPIMulator/golang_vm/uPIMulator/ \
        --bin_dirpath /uPIMulator/golang_vm/uPIMulator/bin/ \
        --benchmark GEMV \
        --logic_frequency 1600 \
        --memory_frequency 3200 \
        --num_channels 8 \
        --num_ranks_per_channel 2  \
        --num_dpus_per_rank "$NUM_DPUS" \
        --num_tasklets 16 \
        --data_prep_params 8192 \
        --t_rcd 32 \
        --t_ras 78 \
        --t_rp 32 \
        --t_cl 32 \
        --t_bl 8

    # Zip the bin directory
    ZIP_FILE="GEMV${NUM_DPUS}.zip"
    echo "Creating archive: ${ZIP_FILE}"
    zip -r "${ZIP_FILE}" bin

    # Clean up bin directory
    echo "Cleaning up bin directory..."
    rm -rf bin
    mkdir bin

    echo "Completed iteration for --num_dpus_per_rank=${NUM_DPUS}"
done

echo "All runs completed!"
